echo "===SeetaFace2==="

git clone https://github.com/seetafaceengine/SeetaFace2.git ./face/SeetaFace2
cd ./face/SeetaFace2 || exit
git checkout -b develop
git branch --set-upstream-to=origin/develop develop
git pull --all
mkdir build
cd build || exit
cmake ..
make && sudo make install

echo "===Finished==="
