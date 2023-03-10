# Тестовое задание для поступления в GoCloudCamp

Результаты выполнения тестового задания следует опубликовать на GitHub и отправить на почту **gocloudcamp@sbercloud.ru** до 23:59:59 10 марта 2023 г. Практическое задание состоит из нескольких частей, при этом, даже если вы сделали только одну часть задания - присылайте ее на рассмотрение. 

## 1. Вопросы для разогрева

1. Опишите самую интересную задачу в программировании, которую вам приходилось решать?

Самая интересная задача в программировании, которую мне приходилось решать, была связана с разработкой алгоритма машинного обучения для автоматической классификации образов кириллических и латинских букв. Такой алгоритм я встроил в своё мобильное приложение. Это было интересно, потому что требовалось применять знания из разных областей, таких как компьютерное зрение, статистика и машинное обучение, чтобы создать эффективную модель классификации. Мне пришлось ознакомиться с большим объёмом информации, которая ранее мне была незнакома. В процессе решения этой задачи я получил ценный опыт в области машинного обучения и приобрел навыки работы с большими объемами данных и сложными алгоритмами.

2. Расскажите о своем самом большом факапе? Что вы предприняли для решения проблемы?

Мой самый большой факап был связан с неправильным размещением файла на продакшн-сервере, что привело к сбою в работе мобильного приложения с тематикой футбола. Я быстро понял, что файл был размещен не в том каталоге, и смог быстро исправить ошибку. Для предотвращения подобных ситуаций в будущем, я улучшил процесс размещения файлов на сервере и создал дополнительные для себя проверки на этапе развертывания приложения.

3. Каковы ваши ожидания от участия в буткемпе?

Я хотел бы погрузиться в промышленную разработку и получить опыт работы в реальном проекте под руководством опытных менторов. Я также надеюсь научиться новым технологиям и подходам к разработке, которые могут быть полезны в будущей работе. Я уверен, что участие в буткемпе даст мне возможность улучшить свои навыки и стать более компетентным разработчиком.

---

## 2. Разработка музыкального плейлиста

### Часть 1. Разработка основного модуля работы с плейлистом

Требуется разработать модуль для обеспечения работы с плейлистом. Модуль должен обладать следующими возможностями:
 - Play - начинает воспроизведение
 - Pause - приостанавливает воспроизведение
 - AddSong - добавляет в конец плейлиста песню
 - Next воспроизвести след песню
 - Prev воспроизвести предыдущую песню

#### Технические требования

 - Должен быть описан четко определенный интерфейс для взаимодействия с плейлистом
 - Плейлист должен быть реализован с использованием двусвязного списка.
 - Каждая песня в плейлисте должна иметь свойство Duration.
 - Воспроизведение песни не должно блокировать методы управления.
 - Метод воспроизведения должен начать воспроизведение с длительностью, ограниченной свойством Duration песни. Воспроизведение должно эмулироваться длительной операцией.
 - Следующая песня должна воспроизводиться автоматически после окончания текущей песни.
 - Метод Pause должен приостановить текущее воспроизведение, и когда воспроизведение вызывается снова, оно должно продолжаться с момента паузы.
 - Метод AddSong должен добавить новую песню в конец списка.
 - Вызов метода Next должен начать воспроизведение следущей песни. Таким образом текущее
 - спроизведение должно быть остановлено и начато воспроизведение следущей песни 
 - Вызов метода Prev должен остановить текущее воспроизведение и начать воспроизведение предыдущей песни.
 - Реализация метода AddSong должна проводиться с учетом одновременного, конкурентного доступа.
 - Следует учитывать, что воспроизведение может быть остановлено извне 
 - Реализация методов Next/Prev должна проводиться с учетом одновременного, конкурентного доступа.
 - Примечание: Все реализации должны быть тщательно протестированы и оптимизированы с точки зрения производительности.

### Часть 2: Построение API для музыкального плейлиста

Реализовать сервис, который позволит управлять музыкальным плейлистом. Доступ к сервису должен осуществляться с помощью API, который имеет возможность выполнять CRUD операции с песнями в плейлисте, а также воспроизводить, приостанавливать, переходить к следующему и предыдущему трекам. Конфигурация может храниться в любом источнике данных, будь то файл на диске, либо база данных. Для удобства интеграции с сервисом может быть реализована клиентская библиотека.

### Технические требования

* реализация задания может быть выполнена на любом языке программирования
* сервис должен обеспечивать персистентность данных
* сервис должен поддерживать все CRUD операции 
* удалять трек допускается только если он не воспроизводится в данный момент
* API должен иметь необходимые методы для взаимодействия с плейлистом.
* API должен возвращать значимые коды ошибок и сообщения в случае ошибок.


### Будет здорово, если:
* в качестве протокола взаимодействия сервиса с клиентами будете использовать gRPC
* напишите Dockerfile и docker-compose.yml
* покроете проект unit-тестами
* сделаете тестовый пример использования написанного сервиса
