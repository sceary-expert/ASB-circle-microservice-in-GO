# social-media-circles-microservice

## POST request for create circle handle

```
model := new(models.CreateCircleModel)
```
model is initialized as an empty create-circle-model.
Now the data from the requst body is passed by using `c.BodyParser`.

If `model.name` is same as `followingCircleName` (`"following"`) then it will throw an error saying that can not use `following` as a circle name. 
```
currentUser, ok := c.Locals(types.UserCtxName).(types.UserContext)
```
This line is used for extracting currentUser.
Now a new circle is created and ownerUSerId of that circle, Name is assigned. IsSystem is initialized to false initially.
```
circleService, serviceErr := service.NewCircleService(database.Db)
```
This line is used for creating a new circle service. 
Let's have a look at how this is created : 

Basically it creates a new data repository 

Now the `circleService` gets saved by calling the `SaveCircle` function.

This `SaveCircle` function call the following channel in order to save the `circleService`.
```
result := <-s.CircleRepo.Save(circleCollectionName, circle)
```
If no error occured we return the json with `objectId` of that context fibre as objjectId of newCircle.