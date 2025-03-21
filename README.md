# Lab2 - Mathematical Expression Converter

## Опис

Lab2 - це програма для конвертації математичних виразів із постфіксної (Reverse Polish Notation, RPN) в інфіксну нотацію. Програма реалізує обробку вхідних математичних виразів, перевірку на помилки та виведення результату у форматі інфіксного виразу.

Програма використовує структуру `ComputeHandler`, яка відповідає за отримання даних з вхідного потоку, їх обробку та виведення результату у вихідний потік.

## Алгоритм

Програма реалізує алгоритм перетворення виразу з постфіксної нотації в інфіксну нотацію. Вхідний вираз спочатку зчитується з потоку, потім конвертується в інфіксну нотацію, а результат виводиться у вихідний потік.

Основні етапи роботи програми:

1. **Читання вхідних даних**: Вхідний вираз зчитується з потоку типу `io.Reader` за допомогою методу `io.Copy` та зберігається у буфер.
2. **Обробка виразу**: Зчитаний вираз очищується від зайвих пробілів за допомогою `strings.TrimSpace`.
3. **Перевірка на порожній рядок**: Якщо після очищення рядок порожній, програма виводить помилку.
4. **Конвертація постфіксного виразу в інфіксний**: Для цього використовується алгоритм (реалізація функції `PostfixToInfix`).
5. **Виведення результату**: Результат конвертації записується у вихідний потік за допомогою `fmt.Fprintln`.

## Інсталяція

### Кроки для запуску:

1. Клонуйте репозиторій:
   ```bash
   git clone https://github.com/DmytroHalai/kpi-2.git
   ```
2. Перейдіть в каталог проєкту:
   ```bash
   cd kpi-2
   ```
3. Встановіть всі необхідні залежності (якщо вони є):
   ```bash
   go mod tidy
   ```
2. Перейдіть в директорію, де розташована точка запуску програми:
   ```bash
   cd cmd\example
   ```
4. Для запуску програми використовуйте команду (про флаги детальніше далі):
   ```bash
   go run main.go <флаг>
   ```

## Використання

Програма дозволяє обробляти математичні вирази у постфіксній нотації як з консолі, так і з файлів. Ви можете передавати вирази для обробки через флаги при запуску програми.

### Запуск програми

Програма підтримує три основні флаги:

- `-e`: для введення постфіксного виразу безпосередньо через консоль.
- `-f`: для вказання файлу, що містить постфіксний вираз.
- `-o`: для виведення результату в файл (за замовчуванням результат виводиться в консоль).

### Флаги та їх використання:

#### Введення та виведення через консоль

Якщо ви хочете ввести постфіксний вираз через консоль і побачити результат виведення також у консоль, використовуйте наступну команду:

```bash
go run main.go -e "3 4 + 5 *"
```

У цьому випадку програма виведе інфіксну форму виразу `((3 + 4) * 5)` в консоль.

#### Введення з файлу та виведення в консоль

Якщо ви хочете отримати вираз з файлу, а результат вивести в консоль, скористайтеся таким флагом:

```bash
go run main.go -f "input.txt"
```

Де `input.txt` містить постфіксний вираз. Результат буде виведений в консоль.

#### Введення та виведення через файли

Щоб і вводити, і виводити через файли, скористайтеся наступною командою:

```bash
go run main.go -f "input.txt" -o "output.txt"
```

У цьому випадку вираз зчитується з файлу `input.txt`, а результат конвертації виводиться у файл `output.txt`.

### Обмеження

- Ви не можете одночасно вказати два джерела вводу. Програма дозволяє лише одне джерело: або `-e` (вираз у вигляді рядка), або `-f` (файл з виразом). Якщо вказано обидва, буде виведено помилку.

### Приклад запуску:

1. **Конвертація виразу через консоль**:
    ```bash
    go run main.go -e "3 4 + 5 *"
    ```
   **Виведення**:
    ```
    ((3 + 4) * 5)
    ```

2. **Конвертація виразу з файлу**:
    ```bash
    go run main.go -f "input.txt"
    ```
   **Виведення**:
    ```
    ((3 + 4) * 5)
    ```

3. **Виведення результату в файл**:
    ```bash
    go run main.go -f "input.txt" -o "output.txt"
    ```

## Тестування

Програма має набір тестів, які перевіряють коректність роботи алгоритму перетворення постфіксного виразу в інфіксну форму.

### Успішні збірки:
Для перевірки успішних збірок та виконання програми зверніться до наступних посилань:
- [Успішні збірки](https://github.com/DmytroHalai/kpi-2/actions/workflows/build.yml?query=is%3Asuccess)

### Неуспішні збірки:
Якщо виникли помилки під час збірки, ви можете перевірити статус невдалих спроб:
- [Неуспішні збірки](https://github.com/DmytroHalai/kpi-2/actions/workflows/build.yml?query=is%3Afailure)
