# social-media-circles-microservice

## POST request for create circle handle

```model := new(models.CreateCircleModel)```
model is initialized as an empty create-circle-model.
Now the data from the requst body is passed by using `c.BodyParser`.