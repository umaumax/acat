# acat

asynchronous multiple read file cat command

```
# terminal A
mkfifo hogepipe
cat hogepipe | gdb
# or
acat - hogepipe | gdb

# terminal B
# you can add rlwrap
cat > hogepipe

# teminal C
echo info > hogepipe
```
