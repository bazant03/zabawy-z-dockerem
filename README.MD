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

# C.D.N.