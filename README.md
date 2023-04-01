## BourbonGo

This is a small REST API written in Golang using a flat-file SQLite database with information
on various types of bourbon. It was written as an example application after finishing a Golang course.

### Components
- Golang
    - [Gin](https://github.com/gin-gonic/gin) 
    - [Squirrel](https://github.com/Masterminds/squirrel)
- SQLite 

### Routes

| **Method** | **URI**           | **Description**                                 |
|------------|-------------------|-------------------------------------------------|
| GET        | `/v1/bourbons`    | List all bourbons.                              |
| GET        | `/v1/bourbon/:id` | Get a specific bourbon by numeric ID.           |
| POST       | `/v1/bourbon`     | Create a new bourbon record.                    |
| PUT        | `/v1/bourbon/:id` | Update a bourbon with the specified numeric ID. |
| DELETE     | `/v1/bourbon/:id` | Delete a bourbon with the specified numeric ID. |

## Model Stucture
```json
{
    "id": 2,
    "name": "A. Smith Bowman Distillery Bowman Brothers Small Batch Straight Bourbon Whiskey",
    "size": "750ml",
    "price": 28.99,
    "abv": 45,
    "description": "John J., Abraham, Joseph, and Isaac Bowman were Virginia militia officers in the American Revolutionary War. This hand-crafted bourbon whiskey is a tribute to their heroism. Our Bowman Brothers Small Batch Bourbon is distilled three times using the finest corn, rye, and malted barley, producing distinct hints of vanilla, spice, and oak."
}