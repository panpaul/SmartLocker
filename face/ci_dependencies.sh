echo "===SeetaFace2==="

git clone https://github.com/seetafaceengine/SeetaFace2.git ./face/SeetaFace2
cd ./face/SeetaFace2 || exit
git checkout -b develop
git branch --set-upstream-to=origin/develop develop
git pull --all
mkdir build
cd build || exit
cmake .. -DCMAKE_INSTALL_PREFIX=./install
make && sudo make install

echo "===Finished==="


echo "===ncnn==="

git clone https://github.com/Tencent/ncnn ./face/ncnn
cd ./face/ncnn || exit
mkdir build
cd build || exit
cmake .. -DCMAKE_INSTALL_PREFIX=./install
make && sudo make install

echo "===Finished==="
