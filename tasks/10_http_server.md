### Завдання: HTTP server з використанням вбудованих пакетів

#### Основні Цілі
Це завдання призначене для того, щоб ви могли практично застосувати навички реалізації HTTP request handler's

#### Завдання

Скопіюйте приклад з `/examples/chapter_10_base/05_resful_simple`
1. Додайте механізм `pagination` в функції `getProducts` для списку товарів. Варіанти реалізації:
   - через параметри (Query)
   - через ключі в заголовку запросу (`r.Header.Get("{ключ в заголовку}")`)
    
    Бажано реалізувати обидва варіанта

2. Зробіть модель данних для замовника товарів та реалізуйте CRUD в роуті `/customers`

   Модель `Customer` данних повинна включати:
   - Идентіфікатор
   - Назву
   - Код 

Список замовників буде повернуто функцією пакету models.GetСustomers().

Список замовників має імпемінтувати фукціонал аналогічний функціоналу `ProductsList`:
  - Get
  - Add
  - Delete
  - Update

Aле "базовим" сховищем для зберігання данних має бути тип map[int]*Customer

**Тестування**

Ваш сервер повинен працювати як з роутом `/products` так і з створенним Вами `/customers` імпементуючи методи HTTP:
 - GET
 - POST
 - PUT
 - DELETE

 Для тестування можна використовувати `Postman` або `Apidog` https://apidog.com/compare/apidog-vs-postman/?utm_source=google_Ia&utm_campaign=20846090569&utm_content=154812052205&utm_term=get%20postman&gclid=CjwKCAiAibeuBhAAEiwAiXBoJDtNmGwTtysgkMALvHWdluA_xZ-29GEDh3xJrut39vngX8W1s-iSaBoCz7gQAvD_BwE
