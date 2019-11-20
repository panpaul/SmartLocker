echo "===SeetaFace2==="

git clone https://github.com/seetafaceengine/SeetaFace2.git ./face/SeetaFace2
cd ./face/SeetaFace2 || exit
git checkout -b develop
git branch --set-upstream-to=origin/develop develop
git pull --all
mkdir build
cd build || exit
cmake .. -DCMAKE_INSTALL_PREFIX=./install
make -j4 && sudo make install

echo "===Finished==="

cd ../../../

echo "===ncnn==="

git clone https://github.com/Tencent/ncnn ./face/ncnn
cd ./face/ncnn || exit
mkdir build
cd build || exit
cmake .. -DCMAKE_INSTALL_PREFIX=./install
make -j4 && sudo make install

echo "===Finished==="

cd ../../../

echo "===json==="

git clone https://github.com/nlohmann/json.git ./face/json
cd ./face/json || exit
mkdir build
cd build || exit
cmake .. -DCMAKE_INSTALL_PREFIX=./install -DJSON_BuildTests=OFF -DBUILD_TESTING=OFF
make -j4 && sudo make install

echo "===Finished==="
