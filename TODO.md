- [ ] Make it at least compilable | 적어도 컴파일은 되게 만들기

  - [ ] Design .rgoc file | .rgoc 파일 설계하기

- [ ] Integrate internal/compiler with internal/RVM | compiler와 RVM 묶기
- [ ] Localize compiler output | 컴파일러 출력 로컬라이징
- [ ] Inspect some logics | 로직 검토
- [ ] Documentation | 문서화
- [ ] Refactor | 리팩토링

  - [ ] Make it cross-compilable | 크로스 컴파일 가능하게 만들기

    - [ ] Separate libraries to each repository by platform and use git submodule | 플랫폼별로 라이브러리를 각각의 저장소로 분리하고 git submodule을 사용하기
    - [x] Change some Windows specific codes to cross-platform | 윈도우 전용 코드를 크로스 플랫폼으로 변경하기

    - [x] Use [external library](https://github.com/hajimehoshi/go-steamworks) instead of internal/steam | internal/steam 대신 [외부 라이브러리](https://github.com/hajimehoshi/go-steamworks) 사용하기

  - [ ] Separate rgo directory to example repository | rgo 디렉토리를 예제 저장소로 분리하기

  - [x] Follow [standard project layout](https://github.com/golang-standards/project-layout) | [표준 프로젝트 레이아웃](https://github.com/golang-standards/project-layout) 따르기

- [ ] Export game to a single executable file | 게임을 내보낼 때 단일 실행 파일로 내보내기
  - [ ] Static compilation of libraries | 라이브러리 정적 컴파일
