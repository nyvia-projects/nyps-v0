# nyps 

An intuitive CLI package manager and utility tool built with Go, designed to streamline the installation and management of scripts and packages through efficient communication with a web server, further enriched with potential file sharing and chat capabilities for enhanced user collaboration and resource sharing.

- **user story 1**: *I would like to share **gpg** or **docker-compose** file with a colleague*
- **user story 2**: *I would like to share a **secret** or sensitive **.env** through a secure channel* 
- **user story 3**: *I would like to install a **script/package** already available in my server repository*
  
#### Use Case 1:
- Title: Getting/Installing a script(let's say a very specific **ansible playbook**)
- Primary Actor: User (IT enthusiast or developer)
- Example Success Scenario:
  1. User runs ```$ nyps get ansible-playbooks/enable-ssh-key-login```
  2. System retrieves desired ***.yaml*** playbook.
  3. System installs(puts) on already preconfigured location by user(ex. ```~/lab/network-ops/playbooks/```)


#### Use Case 2:
- Title: Getting/Installing a Package(let's a custom messaging application - [nychat](https://github.com/nyvia-projects/codename-nc))
- Primary Actor: User (IT enthusiast or developer)
- Example Success Scenario:
  1. User runs ```$ nyps install nychat```
  2. System retrieves desired package from a web server.
  3. System installs package on user's machine.
  4. System confirms successful installation and prints ***help*** command output
    - Exceptions:
        - Package not found
        - Package not available (under dev or deprecated)

## Notes

- Configuration file in ```~/.nypsrc```, where user put various configurations
- Should have or use a secure solution for storing sensitive configuration data (ex. recipients)
- Should work both in local network and public internet.

I am still shaping what I want **nyps** to become, and I will continue updating this as I develop it.