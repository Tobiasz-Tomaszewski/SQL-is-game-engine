As some youtuber once said: 'Anything is a game engine, if you try hard enough'. You can wathc the video I am talking about [here](https://www.youtube.com/watch?v=djIufZ7Fyms).

I am doing this project to further prove this statement. The game engine I've choosen is SQL database. I am trying to learn stuff in a meantime (mainly golang and developing a microservice based solution that can be deployed on a kubernetes cluster).

The basic idea looks like this:

```mermaid
graph LR
    Player --> movement-handler;
    movement-handler --> world-handler;
    world-handler --> game-sqlgine;
    world-viewer --> world-handler;
```
