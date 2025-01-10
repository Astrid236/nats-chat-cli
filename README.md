CHAT CLI CON GO Y NATS
Esta aplicación consiste en un servidor NATS que permite la comunicación entre diferentes usuarios al unirse a un canal de chat. Estos usuarios pueden enviar y recibir mensajes de otros usuarios en tiempo real. Para el desarrollo de la aplicación se ha utilizado GO. 

COMPONENTES: 
  - main.go: código main de la aplicación
  - servidor nats: fichero docker-compose.yml
    
INSTRUCCIONES DE CONFIGURACIÓN:

1. Clonar el repositorio:
   git clone https://github.com/Astrid236/nats-chat-cli.git

2. Iniciar servidor NATS:
   docker compose up

3. Ejecutar la aplicación:
   go run main.go <NATS_SERVER> <CHANNEL> <USERNAME>

   Ejemplo: go run main.go nats://localhost:4222 chat Bob

4. Abrir otro terminal para conectarse al mismo canal.

   Abrir una nueva ventana del terminal y ejecuar la aplicación con un nombre de usuario diferente.

DEPENDENCIAS: 
-  github.com/nats-io/nats.go -> Biblioteca cliente de NATS para conectar al servidor
-  bufio -> para leer entradas desde el terminal
-  os -> para manejar argumentos de la línea de comandos
-  time -> para trabajar con tiempos

DESCONEXIÓN DEL SERVIDOR

Ejecutar: docker compose down

FUNCIONAMIENTO DEL CÓDIGO: 

1. La aplicación toma tres argumentos de entrada, el servidor nats, el nombre del canal y el nombre de usuario.
2. Se levanta el servidor y los clientes se conectan al servidor NATS utilizando la dirección proporcionada como primer argumento.
3. En la aplicación se utilizan los metodos Subscribe y Publish para manejar la comunicación entre los diferentes usuarios conectados al canal. 
   
