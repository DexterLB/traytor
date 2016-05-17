## todo list for traytor

- architecture
    - [ ] move all the code into neat separate packages

- paralellism
    - [ ] try out RPC libraries:
      [rpc](https://golang.org/pkg/net/rpc/),
      [jsonrpc](https://golang.org/pkg/net/rpc/jsonrpc/),
      [gorrilla rpc](http://www.gorillatoolkit.org/pkg/rpc)
      and decide which one is easiest to work with
    - [ ] make the scene serialiseable, so we can call a function that takes
      the scene as an argument remotely (maybe base64 encode the scene
      and pass a single string?)
    - [ ] create a "worker" command which acts as an RPC server, and
      has `uploadScene(scene)` and `renderSample() Image` RPC methods
    - [ ] create a "client" command which tells multiple workers to render
      samples and adds their results together, forming an image
    - [ ] multi-threaded workers - make workers accept several requests at
      once, with a limit on active renders (so that a worker with 4
      processors can accept 8 requests at once, render 4 at once
      and start rendering the next 4 while the first 4 are being sent
      back to the client). This should be much easier than it sounds,
      because buffered channels are awesome.
- GUI
    - [ ] try out [QML](http://doc.qt.io/qt-5/qtqml-index.html) without
      Go - for example use [qmlscene](http://doc.qt.io/qt-5/qtquick-qmlscene.html)
    - [ ] make the same thing run under Go with [go-qml](https://github.com/go-qml/qml)
    - [ ] find a way to display our Image in the GUI (QPainter still not
      available in Go? Maybe need to use QLabel with a custom image? Or OpenGL?)
    - [ ] huge RENDER button!
    - [ ] implement the client in the GUI so that it can display a new image
      each time a sample is received
    - [ ] make a fancy worker selector in the GUI which lists the IP addresses
      of all available workers and the user can select which workers to
      render on
    - [ ] make a stormtrooper logo (maybe render it in Traytor?)

- rendering
    - [ ] fix the goddamn refraction
    - [ ] add mix shader/add shader
    - [ ] add a fresnel sampler
    - [ ] implement lamp sampling or bidirectional path tracing to speed
      it up a lot (hard)
    - [ ] implement matte reflection and refraction
      (very hard, requires statistics knowledge)
    - [ ] add sampler addition, multiplication, screen etc