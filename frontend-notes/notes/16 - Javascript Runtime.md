---
title: 16 - Javascript Runtime
created: '2019-04-03T09:30:02.327Z'
modified: '2019-04-03T12:07:49.402Z'
tags: [JS]
---

# 16 - Javascript Runtime


![](https://raw.githubusercontent.com/8483/notes/master/pics/js_event_loop.png)

The Javascript runtime only knows of the heap and call stack.

The rest of the functionality, like async stuff, is provided in the form of WebAPIs by the browser/Node.

For async operations, the callback is offloaded into a separate queue, which is emptied by the event loop one by one as soon as the stack becomes empty.

Without the event loop the stack would be blocked during the whole duration of the async operaiton, basically freezing the app.



