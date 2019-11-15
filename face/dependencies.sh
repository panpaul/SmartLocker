echo "===SeetaFace2==="

git clone https://github.com/seetafaceengine/SeetaFace2.git ./face/SeetaFace2
cd ./face/SeetaFace2
git pull --all
git checkout -b develop
mkdir build
cd build
cmake .. -DCMAKE_INSTALL_PREFIX=/usr/local
make && make install

echo "===Finished==="