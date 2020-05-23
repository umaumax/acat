# acat

asynchronous multiple read file cat command

## how to use
```sh
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

## FYI
if you don't use `/dev/stdin`, you can use `xargs` with `-P` option
```
mkfifo hogefifo
mkfifo piyofifo
{ echo "cat hogefifo"; echo "cat piyofifo"; } | xargs -L1 -I{} -P2 bash -c "{}"
```

## NOTE
* `io.MultiReader` read synchronously
