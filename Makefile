include .env
export 



export PROJECT_ROOT =${shell pwd}
env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down todoapp-postgres


env-cleanup:
	read -p "Clear all environment volume files? Risk of memory loss. [y/N]: " ans; \
	if [ "$$ans" = "y"  ]; then \ 
		docker compose down todo-app-posgres && \
		rm -rf out/pgdata && \
		echo "Enviroment files cleared"; \
	else \
		echo "Enviroment clearing cancelled"
	fi