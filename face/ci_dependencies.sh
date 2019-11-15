echo "===SeetaFace2==="

git clone https://github.com/seetafaceengine/SeetaFace2.git ./face/SeetaFace2
cd ./face/SeetaFace2
git pull --all
git checkout -b develop
git pull --all
mkdir build
cd build
cmake ..
make && sudo make install

echo "===Finished==="