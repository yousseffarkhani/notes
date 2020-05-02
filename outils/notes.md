console: Interface pour le shell
terminal: Endpoint de communication avec l'utilisateur (représente la fenêtre)
Wrapper autour d'un shell
console / terminal => Même chose
Iterm2, terminal

shell: Programme permettant de lancer d'autres programmes
Représente le CLI
- Interprète les commandes (cd, ls, ...) et les traduit en des appels au kernel 
- Historique de commandes
Il existe plusieurs implémetations spécifiques de shell: bash, zsh, sh, PowerShell

Un terminal run un shell automatiquement.
Un shell peut être lancé sans terminal
Le terminal a été crée pour que l'on puisse interagir avec des shells mais les shells peuvent être lancés en background sans interaction humaine (cron job)

stdin, stdou, stderr: Standard streams sont redirigées vers le terminal

Kernels
Partie de l'OS
Core of the OS
Permet d'interagir avec le hardware, file I/O, créer des process
Handle multi tasking
Manage system ressource
Tasks:
- Memory management
- Shared resource access
- IO access
- Task and job schedulig
- Thread and process management
- Process flow control

Process ?

echo $0 pour connaitre le shell utilisé
