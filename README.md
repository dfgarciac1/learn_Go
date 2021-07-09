# Arquitectura Recomendada para proyectos

Se recomienda una Estructura hexagonal, es una arquitectura atomica  cada parte tiene una interfas o ENDPOINT 
## Comenzando 🚀

_Se puede dar conocimiento a la arquitecuta hexagonal como de puertos y adaptadores, como motivación principal es la separación de capas 
desde la logica de negocio  hasta el framework de trabajo permitiendo independencia entre los Endpoint( representa un método único en la interfaz de nuestro servicio.).

_Cada lado del hexagono representa un puerto el cual se encarga de una tarea en especifico :

### 1)Puerto: definición de una interfaz pública.
### 2)Adapter: especialización de un puerto para un contexto concreto.


![image](https://user-images.githubusercontent.com/53051443/125107995-0bf99b80-e0a7-11eb-8f21-90347c2a3a01.png)


Mira **https://www.encora.com/es/blog/construyendo-microservicios-con-golang-y-go-kit** para conocer como trabajar más sencillo con microservicios y una explicación de este mismo.


