- Envoyer un io.Writer vers un io.Reader
Afficher le résultat d'un http.Get()
```go
io.Copy(os.Stdout, resp.Body)
```