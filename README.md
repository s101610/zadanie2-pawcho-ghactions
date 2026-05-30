## Zadanie 2 – GitHub Actions, GHCR, cache DockerHub i test CVE

### Aplikacja testowa

W projekcie znajduje się prosta aplikacja napisana w Go. Uruchamia ona serwer HTTP na porcie `8080` i udostępnia dwa endpointy:

```text
/
/health
```

Endpoint `/` zwraca odpowiedź w formacie JSON, a `/health` służy do sprawdzenia, czy aplikacja działa poprawnie.

Aplikację lokalnie można uruchomić poleceniem:

```bash
go run ./src
```

Po uruchomieniu aplikacja jest dostępna pod adresami:

```text
http://localhost:8080
http://localhost:8080/health
```

---

### Dockerfile

Do budowy obrazu użyto budowy wieloetapowej. W pierwszym etapie aplikacja jest kompilowana w obrazie `golang`, a drugi etap bazuje na `scratch`, dzięki czemu obraz końcowy jest mały i zawiera tylko plik wykonywalny.

Ważny fragment konfiguracji:

```dockerfile
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /dist/server ./src
```

Zmienne `TARGETOS` i `TARGETARCH` są przekazywane przez BuildKit, dlatego ten sam `Dockerfile` może posłużyć do budowy obrazu dla `amd64` i `arm64`.

---

### Podsumowanie

Przygotowane rozwiązanie spełnia wymagania zadania. Workflow buduje obraz dla architektur `linux/amd64` i `linux/arm64`, korzysta z cache w DockerHub, wykonuje skan CVE dla podatności `HIGH` i `CRITICAL`, a następnie publikuje obraz do GHCR tylko wtedy, gdy test bezpieczeństwa zakończy się poprawnie. Obraz jest tagowany na podstawie commita lub tagów SemVer, bez używania niejednoznacznego tagu `latest`.
