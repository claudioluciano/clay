package resource

import (
	"embed"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"path"
	"reflect"
	"strings"
)

// resourceHandler is an internal intermediate struct which stores the prefix, extension and handler types for the loader.
type resourceHandler struct {
	extension string
	prefix    string
	handler   ResourceHandler
}

// loader is the internal struct responsible for actually loading resources into memory.
type loader struct {
	handlers  []resourceHandler
	resources map[Path]LoadedResource
}

func (l *loader) handleResourceType(prefix string, extension string, handler ResourceHandler) {
	l.handlers = append(l.handlers, resourceHandler{
		extension: extension,
		prefix:    prefix,
		handler:   handler,
	})
}

func (l *loader) loadFromFs(directory string, fs embed.FS) []error {
	if l.resources == nil {
		l.resources = make(map[Path]LoadedResource)
	}

	log.WithField("directory", directory).Trace("Loading from directory")
	var errs []error

	dir, readErr := fs.ReadDir(directory)
	if readErr != nil {
		errs = append(errs, fmt.Errorf("could not read directory %s: %w", directory, readErr))
		return errs
	}

	for _, entry := range dir {
		// Recursive folder includes
		if entry.IsDir() {
			subErrs := l.loadFromFs(path.Join(directory, entry.Name()), fs)
			if len(subErrs) > 0 {
				errs = append(errs, subErrs...)
			}
			continue
		}

		// Files in directory
		handleFileErr := l.handleFileFromDir(entry, directory, fs)
		if handleFileErr != nil {
			errs = append(errs, handleFileErr)
		}
	}

	return errs
}

func (l *loader) addResource(path Path, instance any) error {
	existing, exists := l.resources[path]
	if exists {
		return fmt.Errorf("can not add resource at path %s because it already exists (typeof %s)", path, reflect.TypeOf(existing))
	}

	l.resources[path] = LoadedResource{
		Path:     path,
		instance: instance,
	}
	return nil
}

func (l *loader) handleFileFromDir(entry fs.DirEntry, directory string, fs embed.FS) error {
	loadPath := path.Join(directory, entry.Name())
	file, readFileErr := fs.Open(loadPath)
	if readFileErr != nil {
		return fmt.Errorf("unable to read file at path %s: %w", loadPath, readFileErr)
	}
	defer file.Close()

	var handled bool
	for _, h := range l.handlers {
		// Skips non-handled ones
		if !strings.HasSuffix(entry.Name(), h.extension) {
			continue
		}

		instance, handlerErr := h.handler.Load(file)
		if handlerErr != nil {
			return fmt.Errorf("handler %s returned error for file path %s: %w", reflect.TypeOf(h.handler), loadPath, handlerErr)
		}
		name := strings.Split(entry.Name(), ".")
		prefixedPath := Path(fmt.Sprintf("%s:%s", h.prefix, name[0]))

		addResourceErr := l.addResource(prefixedPath, instance)
		if addResourceErr != nil {
			return fmt.Errorf("can not add resource: %w", addResourceErr)
		}
		handled = true
		log.Tracef("Added resource %s of type %s", prefixedPath, reflect.TypeOf(instance))
	}

	if !handled {
		return fmt.Errorf("file type was not handled during loading for path %s", loadPath)
	}

	return nil
}
