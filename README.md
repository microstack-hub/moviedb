# moviedb

## Exercise - 🎥🎥🎥
To demonstrate how the ambassador pattern works, we will use [The Movie DB](https://www.themoviedb.org/). Head over to the website and register (it's free) to get an API key.

The Movie DB website offers a REST API where you can get information about the movies. 

first create the secret object as follows using the API key

```bash
kubectl create secret generic themoviedb --from-literal=apikey=<YOUR API KEY>
```
output:
```bash
secret/themoviedb created
```

and apply the file
```bash
kubectl apply -f moviedb.yaml`

Use the exec command to run the curl command inside the main container:

```bash
kubectl exec -it themoviedb -c main -- curl localhost:8080/movies
```
What legendery movies can you spot in the response? ... or not?
