# dothub - Dotfile manager

## Usage

Calling dothub will start the configuration wizard. Every command line argument can be used in the configuration file or as an option id the CLI.


Start monitoring dotfiles from github:

```bash
dothub init --repository username/reponame
branch (hostname): 
```

Start monitoring dotfiles from any git repo, creating a new branch named test:

```bash
dothub init --repository https://hostname/reponame.git --branch test
```

Start monitoring files from a different branch (not master):

```bash
dothub init --repository --origin mybranch username/reponame
branch (hostname):
```

Keeping files updated:

```bash
dothub update
```

Add a file (or directory) from:

```bash
dothub add file.txt
```

Remove file from monitoring (does not delete files):

```bash
dothub remove file.txt
```

## Configuration

```bash
# ~/.dothubrc
repository: https://github.com/username/reponame
branch: $hostname
# origin: master
# message: $user@$hostname at $date $time changed: $files
monitoring:
  - .bashrc
  - .bash_aliases
  - .gitignore
  - .profile
  - .npmrc
  - .tmux.conf
  - .vimrc
  - bin
  - .yarnrc
#/etc/dothubrc
# config: ~/.dothub
```


