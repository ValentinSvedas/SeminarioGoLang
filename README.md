# SeminarioGoLang

### Consigna

Crear una funcion que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena. Por ejemplo:

TX03ABC

**Para esto implemento un programa para la detección de cadenas de caracteres con determinada estructura, en donde cada una debe de respetar un estandar por lo cual se creó un metodo en donde compruebe que esa cadena es aceptada. Todo fue implementado mediante el lenguaje de programación GoLang**

#### Resolución 
Se creo la clase "Chain" con el metodo NewChain para trasformar un string que viene por parametro a una cadena por ejemplo TX03ABC, "TX" es el tipo, "03" el largo y lo que resta el valor "ABC". Además agregué una interfaz y una función que implemente esa interfaz logrando así poder crear una funcion dentro de la interfaz, que verifique que la cadena pasada por parametro cumpla los estandares.

![image](https://user-images.githubusercontent.com/39970358/133549827-6eb8642a-9fd5-4165-ba58-b4fff63c045d.png)



![image](https://user-images.githubusercontent.com/39970358/133549779-a08a0150-87ec-4beb-b7d7-89d09f212dd6.png)


![image](https://user-images.githubusercontent.com/39970358/133551987-94da43ec-b383-427b-b8a7-8ef1542bf665.png)


![image](https://user-images.githubusercontent.com/39970358/133551602-9dd4358b-8ade-404b-ac83-12a28499d66b.png)


## Consideraciones 👍
- [x] Agregar test de unidad
- [x] Se espera una cobertura del 80% o mas
- [x] Tener en cuenta distintos casos para los test, por ejemplo cadenas inválidas.
- [x] Usar go modules (go mod)
- [x] Entregar un repositorio de github publico con el README.md que explique la solucion que hicieron
