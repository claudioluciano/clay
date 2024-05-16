package resource

import "io"

// ResourceHandler represents a type that handles a specific file extension for loading.
type ResourceHandler interface {
	Load(reader io.ReadCloser) (any, error)
}

// resourceHandler is an internal intermediate struct which stores the prefix, extension and handler types for the loader.
type resourceHandler struct {
	extension string
	prefix    string
	handler   ResourceHandler
}

// loader is the internal struct responsible for actually loading resources into memory.
type loader struct {
	handlers []resourceHandler
}

func (l *loader) handleResourceType(prefix string, extension string, handler ResourceHandler) {
	l.handlers = append(l.handlers, resourceHandler{
		extension: extension,
		prefix:    prefix,
		handler:   handler,
	})
}
