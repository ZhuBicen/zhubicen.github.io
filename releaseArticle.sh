cd docs || exit 1
git pull origin master || exit 1
cd ..
./bin/hugo || exit

go run ./bin/HugoUtils.go || exit 1       

cd docs
git add *
git commit -m"New Article" -a
git push origin HEAD:master
