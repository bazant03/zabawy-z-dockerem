# Kurs Docker dla koła naukowego SOIP

Repozytorium zawiera materiały szkoleniowe oraz przykłady wykorzystania technologii Docker. Znajdziesz tutaj zarówno proste przykłady, jak i bardziej złożone konfiguracje aplikacji.

## 1. Czym jest Docker?
Na początek, aby wyrobić sobię intuicję, możemy porównać Dockera do narzędzi takich jak Virtualbox lub VMware.
Choć idea jest podobna, to jednak wykorzystywane mechanizmy i samo zastosowanie jest bynajmniej nieco inne. 
Docker to narzędzie realizujące koncepcję **wirtualizacji lekkiej**, czyli **konteneryzacji**. W odróżnieniu od klasycznych maszyn wirtualnych, kontenery współdzielą jądro (kernel) systemu operacyjnego maszyny hosta, tzn. procesy uruchomione w kontenerze pracują bezpośrednio w systemie operacyjnym hosta razem ze "zwykłymi" procesami. 

Architektura środowiska opiera się na modelu klient-serwer:
*   **Klienta (CLI):** Interfejsu, za pomocą którego wydaje się komendy.
*   **Serwera (Daemon):** Silnika działającego w tle, który zarządza kontenerami.

## 2. Jak pobrać Dockera?
Na komputerach z systemem Linux, dla większości dystrybucji można go zainstalować z użyciem menadżera pakietów danej dystrybucji.
Najprostszym sposobem instalacji na systemach Windows i macOS jest pobranie **Docker Desktop**.

Instrukcje dostępne są w oficjalnej dokumentacji:
*   Docker Engine - Linux: https://docs.docker.com/engine/install/
*   Docker Desktop - Windows, macOS, (również Linux): https://docs.docker.com/desktop/

## 3. Czym są obrazy?
Obraz to ustandaryzowany pakiet, w którym znajdują się binarki, biblioteki, pliki konfiguracyjne itp. Jednym słowem wszystkie zasoby których będzie potrzebował kontener do swojego działania. 

Dwie istotne cechy obrazów:
1. Niemutowalność, po uruchomieniu kontenera możemy zmodyfikować jego zawartość (kontenera), jednak nie wpłynie to na zawartość samego obrazu na podstawie którego został utworzony kontener. 
2. Obraz jest zbudowany z warstw. Pojedyncza warstwa to zmiana w systemie plików, t.j. dodanie, modyfikacja lub usunięcie pliku.

Tutaj można przeczytać trochę więcej:
*   https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/
*   https://martinheinz.dev/blog/44

Obrazy pobierane są z rejestru obrazów. Domyślnym rejestrem jest *[DockerHub](https://hub.docker.com/)*.
Swoje rejestry posiadają m.in. duże organizacje jak np. GitHub (GHCR) i Oracle (OCR). 
Możemy także utworzyć swój prywatny rejestr, ale to jeszcze nie teraz.

Do zarządzania obrazami służy polecenie **docker image**. Znaczna część tych poleceń ma swoją skróconą formę, co jest charakterystyczne dla CLI Dockera. 

### Przykłady wykorzystania:

* **Pobieranie obrazu z rejestru.**
    W poleceniu podajemy nazwę rejestru, nazwę obrazu (repozytorium) oraz tag (wersję obrazu):
    ```bash
    docker image pull docker.io/ubuntu:latest
    ```
    W tym poleceniu możemy pominąć słowo kluczowe image, dodatkowo jeśli korzystamy z rejestru DockerHub i tagu (`latest`), możemy pominąć te wartości:
    ```bash
    docker pull ubuntu
    ```

* **Wyświetlanie pobranych obrazów.**
    Pokazuje listę obrazów dostępnych lokalnie na maszynie wraz z ich rozmiarami i identyfikatorami.
    ```bash
    docker image ls
    ```
    Wersja skrócona:
    ```bash
    docker images
    ```

* **Usuwanie obrazu.**
    Pozwala usunąć lokalny obraz, podając jego nazwę (unikalną, ew. z rejestrem i tagiem) lub ID. Obraz nie może być używany przez żaden (nawet zatrzymany) kontener.
    ```bash
    docker image rm alpine:3.24.1
    ```
    Wersja skrócona:
    ```bash
    docker rmi alpine:3.24.1
    ```
## 4. Zarządzanie kontenerami
Schemat przedstawiający cykl życia kontenera:

```text
                  +-------------------------+
                  |        [ Obraz ]        |
                  +-------------------------+
                       |                 |
                       |                 | docker create
                       |                 v
                       |           +-------------------+
                       |           |      Stopped      |
                       |           +-------------------+
                       |                 |
                       | docker run      | docker start
                       |                 |  
                       v                 v      docker pause 
            +---- +-------------------------+ ---------------> +-------------------+
            |     |         Running         |                  |      Paused       |
            |     +-------------------------+ <--------------- +-------------------+
            |        |      |             ^    docker unpause
    --rm    |        |      |             |
            |        | *    | docker stop | docker start
            |        |      | docker kill |
            |        v      v             |
            |     +-------------------------+
            |     |         Stopped         |
            |     +-------------------------+
            |          |
            |          | docker rm
            v          v
            +---> +-------------------------+
                  |         Deleted         |
                  +-------------------------+
```

# C.D.N.