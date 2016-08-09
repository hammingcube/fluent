lang=$1

mkdir -p $lang

sed -n -e s/cpp/$lang/g -e "w $lang/gen.go" cpp/gen.go