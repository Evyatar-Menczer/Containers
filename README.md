# Containers
In this project, I manage Docker containers from main server. The main server creates 2 agents in different servers. Each agent gets its amount of containers to create.  
The container's name, image name and amount of containers to create, coming from YAML file. In the main loop, type your commands to control the agents.

**Before first run** make sure to change your permissions to the Docker socket. ```sudo chmod 666 /var/run/docker.sock``` works fine.

# Commands
There are three commands available now:  
```create <PATH>``` - creates the agents, from the YAML file.  
``` delete <NAME>``` - deletes all running ```NAME``` containers.  
```exit now``` - exits the program.
