# ugadunBot
Бот для обучения карточкам

## Цель 
Бот предназанчен для изучения карточек. 

## Входные данные
Карточки хранятся в гугл таблице в следующем формате: 
* `num` -  Номер по порядку. 
* `name` - Название картинки
* `picture` - URL картинки. 

Бот будет получать картинке в формате JSON по запросу

```
curl --location --request GET 'https://script.google.com/macros/s/<tbTocken>/exec'
```

## Логика работы

При получении команды /start бот прогружает таблицу.
Работа ведется в цикле
1. Показывает случайную картинку из поля `picture` и кноппку `"Показать ответ"`
2. После нажания на кнопку `"Показать ответ"` бот выдает название картинки из поля `name` и кнопка `"Следующая"`
3. После нажатия на кнопку `"Следующая"` возвращаемся к п.1

По команде `/reload` бот обновляет картинки из базы и переходит к п.1 цикла

## План
### 1 Делаем каркас.
1. следим за аптедейтами. 
2. отсылаем два сообщения с кнопкой.
    * Картинка с кнопокой `"Показать ответ"`
    * Название картинки с кнопкой ``"Следующая"``

Первоначально реализуем команды:
1. /start   - старт бота
2. /help    - чет тут пишем
3. /next    - следущая картинка. Кнопка `"Следующая"`
4. /answer  - показать ответ. Кнопка `"Показать ответ"`
4. /reload  - загурзить таблицу с нуля. 

Которые просто отправляют сообщения в чат.

ГОТОВО 
### 2 Реализовываем логику работы цикла

ГОТОВО
### 3 Полдключаем картинки
1. Делаем загрузчик картинок
2. 