# Clay
This project started as a way to make golang gamedev easier for myself.

## Built-in components
Clay implements some default components and systems you can use in your game to get a quicker start
with Donburi, and is designed to be flexible.
Any of these can be added to a Donburi entity and will have relevant behaviours.
These will have a default system implemented in the `DefaultPlugins` list.
You can find most of the relevant components in pkg/components. The default systems that use them
are in plugins/render, plugins/audio and plugins/resources.

## Resource loading
Check out the resource package. By default it has support for some common formats.

## WIP
This library is a work in progress. Check out the _example folder.
Basically the actual usage is:
- Your app runs clay
- You register plugins into the instance
- Hook systems onto the plugin

## Contributions
I'm thankful for any contributions and help to make this library better.

