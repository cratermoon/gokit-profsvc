SERVER = cmd/profilesrv
CLIENT = cmd/profilecli


all: client server

client:
	$(MAKE) -C $(CLIENT) build


server:
	$(MAKE) -C $(SERVER) build

run:
	$(MAKE) -C $(SERVER) run

.PHONY: all client server run
