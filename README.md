# Sublime Sync

CLI utility for syncing Sublime Text settings across machines, automagically.

Given a path to your ST3 settings files, and a path to a local .git repo, this tool
will constantly watch for settings changes, package changes, etc. When a `WRITE` event is noticed,
it copies the new file to the local git repo and pushes to Github. That's literally all it does, nothing fancy here. 


## How to Use

### Authentication

This program needs some pre-defined env variables to authenticate itself when syncing a change to Github.

```
GH_USER
GH_PASS
```

Are your user and pass respectively. 

### Program Arguments

This program requires two arguments,

```
--subl      Path to sublime text settings files. On my Ubuntu machine it's in 
            ~/.config/sublime-text-3/Packages/User/ but I'm really not sure
            about Windows or Mac.

--git       Path to directory containing a .git folder. This repo should already have a 
            remote configured otherwise the tool will start freaking out (I didn't code any error handlers for this, so the outcome is... undefined).
```

### Running the Program
```
> go run main.go --git=/path/to/local/repo --subl=/path/to/subl/settings/ 
```
The paths can be relative, but I recommend providing absolute paths here.

Thanks
