# SeminarioGoLang

### Consigna

Crear una función que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena. Por ejemplo:
```
TX03ABC
```
**Para esto implementó un programa para la detección de cadenas de caracteres con determinada estructura, en donde cada una debe de respetar un estándar por lo cual se creó un método en donde compruebe que esa cadena es aceptada. Todo fue implementado mediante el lenguaje de programación GoLang**

#### Resolución 
- Se creó la clase *Chain* con el método *NewChain* para transformar un string que viene por parámetro a una cadena por ejemplo TX03ABC, "TX" es el tipo, "03" el largo y lo que resta el valor "ABC". Además agregué una interfaz y una función que implemente esa interfaz logrando así poder crear una función dentro de la interfaz, que verifique que la cadena pasada por parámetro cumpla los estándares.

- Para tener un test que verifique la función implementada, se añadió *chain_test.go* que comprueba cada situación para verificar que el test pase exitosamente comprobando cada cadena. Lo que hace es instanciar la interfaz mediante la función y crear una cadena según el string que se le pasa, luego con la función *ConfirmChain* devuelve un boolean indicando si la cadena, pasada por parámetro, es aceptada o no.

![image](https://user-images.githubusercontent.com/39970358/133549827-6eb8642a-9fd5-4165-ba58-b4fff63c045d.png)

*(Test pasado por chain.go)*
- Para lograr que se muestre por donde pasó el test, desde la terminal ejecute *go test -cover -coverprofile chtml.out* que setea el nombre en chtml.out y con *go tool cover -html chtml.out* me crea el html para poder ver por donde pasó todo el test marcando en verde todo lo que se testeo correctamente.

- Con el comando *go test ./... -coverprofile=out.test* Se comprobó que el test creado cubra el 91.7% de chain.go.
![image](https://user-images.githubusercontent.com/39970358/133549779-a08a0150-87ec-4beb-b7d7-89d09f212dd6.png)

- Se usó el comando *go mod init github.com/valentinsvedas/tpGoLang* para crear un módulo así poder importar toda la funcionalidad
![image](https://user-images.githubusercontent.com/39970358/133551987-94da43ec-b383-427b-b8a7-8ef1542bf665.png)

- Además en el main cree un mini programa para poder insertar cadenas, con un bucle que se repite si el número insertado es 0. Tuve un inconveniente ya que no me imprime si es verdadera o falsa la cadena verificando la validez del formato.

![image](https://user-images.githubusercontent.com/39970358/133551602-9dd4358b-8ade-404b-ac83-12a28499d66b.png)


## Consideraciones 👍
- [x] Agregar test de unidad
- [x] Se espera una cobertura del 80% o mas
- [x] Tener en cuenta distintos casos para los test, por ejemplo cadenas inválidas.
- [x] Usar go modules (go mod)
- [x] Entregar un repositorio de github publico con el README.md que explique la solucion que hicieron
