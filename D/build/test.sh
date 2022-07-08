#!/usr/bin/env bash

UNZIP=unzip
SRC_FIlE=main.go
BIN_FILE=do
TMP_DIR=tmp
TEST_DIR=tests
RESULT_DIR=result
TEST_FILE=tests.zip

cd ..

# Проверим файл с тестами
if [ ! -e $TEST_FILE ]
then
 echo "Нет файла с тестами"
 exit 1
fi

mkdir -p  $TMP_DIR
mkdir -p $TMP_DIR/$RESULT_DIR

# Есть unzip? Разархивируем
which $UNZIP 1> /dev/null
if [ ! $? -eq 0 ]
then
 echo "Нет $UNZIP"
 exit 1
fi
$UNZIP -o $TEST_FILE  -d $TMP_DIR >/dev/null 2>&1
if [ ! $? -eq 0 ]
then
 echo "Ошибка разархивирования"
 exit 1
fi

# Есть бинарный файл? Компилируем если нет
if [ ! -e $TMP_DIR/$BIN_FILE ]
then
  echo "Компилируем бинарный файл"
  go mod init test
  go mod tidy
  go build -o $TMP_DIR/$BIN_FILE $SRC_FILE
fi


cd  $TMP_DIR/$TEST_DIR


# Убираем ^M в концах файлов теста
in_out_files=$(find .  -maxdepth 1 -regextype sed -regex  "^./*[0-9]\{2\}\(\.a\)\?")
$(sed -i  's/\r$//g' $in_out_files)


# Выбираем входные файлы 
in_files=$(find . -maxdepth 1 -regextype sed -regex  "^./*[0-9]\{2\}" | sed 's/\.\///')

# Запускаем тесты
cd ..
if [ $? -eq 0 ]
then
  for f in  $in_files
  do
    echo "Создаю результаты теста для $f"
    $(cat $TEST_DIR/$f | ./$BIN_FILE > $RESULT_DIR/$f.a)
    if [ ! $? -eq 0 ]
    then
      echo "Ошибка при запуске теста"
      exit 1
    fi
  done
else
  echo "Входные файлы не найдены"
fi

# Сравниваем файлы с результатами тестов
TEST_FAILED=""
for f in  $in_files
do
  diff $TEST_DIR/$f.a $RESULT_DIR/$f.a
  if [ $? -eq 0 ]
  then
    echo "Тест $f прошёл"
  else
    echo "Тест $f провален"
    $TEST_FAILED="YES"
  fi
done

if [ -z "$TEST_FAILED" ]; then
  echo ";-)\nВсе тесты пролшли!\n Удаляю временную папку."
  # Удаляем временную папку
  cd ..
  rm -dfr $TMP_DIR
else
  echo ";-(\n!!!Не все тесты пролшли!!!\n Сморти во временной папке."
fi

