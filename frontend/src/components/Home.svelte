<script>
  import { Link } from "svelte-navigator";
</script>

<div class="px-4 py-5 my-5 text-center">
  <!-- <img class="d-block mx-auto mb-4" src="ksrg-connect_logo.svg" alt="" width="80" height="80"> -->
  <h1 class="display-5 fw-bold">KSRG Connect</h1>
  <div class="col-lg-6 mx-auto">
    <p class="lead mb-4">
      Schnelle Liebe 👀 und Bekanntschaften mit der Mobile-First Website
      KSRG-Connect, die deutschlandweit einzige innovative Website unter 10.000.
      Unterstützt und bereitgestellt von der exklusiven Community des
      Karl-Schmidt-Rottluff Gymnasiums Chemnitz.
    </p>
    <div class="d-grid gap-2 d-sm-flex justify-content-sm-center">
      <!-- <button type="button" class="btn btn-primary btn-lg px-4 gap-3"
        >Primary button</button
      > -->
      <Link
        to="/signup"
        primary={false}
        role="button"
        class="btn btn-primary btn-lg">Get started!</Link
      >
      <!-- <button type="button" class="btn btn-outline-secondary btn-lg px-4"
        >Secondary</button
      > -->
      <Link
        to="/signin"
        primary={false}
        role="button"
        class="btn btn-outline-secondary btn-lg px-4">Sign in</Link
      >
    </div>
  </div>
</div>

<div class="container px-4" id="icon-grid">
  <h2 class="pb-2 border-bottom">Features</h2>

  <div
    class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5"
  >
    <div class="col d-flex align-items-start">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="1.75em"
        height="1.75em"
        fill="currentColor"
        class="bi text-muted flex-shrink-0 me-3"
        viewBox="0 0 16 16"
      >
        <path
          d="M7.657 6.247c.11-.33.576-.33.686 0l.645 1.937a2.89 2.89 0 0 0 1.829 1.828l1.936.645c.33.11.33.576 0 .686l-1.937.645a2.89 2.89 0 0 0-1.828 1.829l-.645 1.936a.361.361 0 0 1-.686 0l-.645-1.937a2.89 2.89 0 0 0-1.828-1.828l-1.937-.645a.361.361 0 0 1 0-.686l1.937-.645a2.89 2.89 0 0 0 1.828-1.828l.645-1.937zM3.794 1.148a.217.217 0 0 1 .412 0l.387 1.162c.173.518.579.924 1.097 1.097l1.162.387a.217.217 0 0 1 0 .412l-1.162.387A1.734 1.734 0 0 0 4.593 5.69l-.387 1.162a.217.217 0 0 1-.412 0L3.407 5.69A1.734 1.734 0 0 0 2.31 4.593l-1.162-.387a.217.217 0 0 1 0-.412l1.162-.387A1.734 1.734 0 0 0 3.407 2.31l.387-1.162zM10.863.099a.145.145 0 0 1 .274 0l.258.774c.115.346.386.617.732.732l.774.258a.145.145 0 0 1 0 .274l-.774.258a1.156 1.156 0 0 0-.732.732l-.258.774a.145.145 0 0 1-.274 0l-.258-.774a1.156 1.156 0 0 0-.732-.732L9.1 2.137a.145.145 0 0 1 0-.274l.774-.258c.346-.115.617-.386.732-.732L10.863.1z"
        />
      </svg>
      <div>
        <h4 class="fw-bold mb-0">Exklusiv</h4>
        <p>
          Durch eine Verfikation wird sichergestellt, dass Personen der
          Community Zugang zu exklusiven Bereichen haben. 💙
        </p>
      </div>
    </div>
    <div class="col d-flex align-items-start">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="1.75em"
        height="1.75em"
        fill="currentColor"
        class="bi text-muted flex-shrink-0 me-3"
        viewBox="0 0 16 16"
      >
        <path
          fill-rule="evenodd"
          d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314z"
        />
      </svg>
      <div>
        <h4 class="fw-bold mb-0">Unterstützung</h4>
        <p>
          Ksrg-Connect ist zu 100% Open Source. Deswegen kannst du diese Website
          aktiv mitentwickeln.
        </p>
      </div>
    </div>
    <div class="col d-flex align-items-start">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="1.75em"
        height="1.75em"
        fill="currentColor"
        class="bi text-muted flex-shrink-0 me-3"
        viewBox="0 0 16 16"
      >
        <path
          d="M6.5 6a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-3z"
        />
        <path
          d="M5.5.5a.5.5 0 0 0-1 0V2A2.5 2.5 0 0 0 2 4.5H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2A2.5 2.5 0 0 0 4.5 14v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14a2.5 2.5 0 0 0 2.5-2.5h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14A2.5 2.5 0 0 0 11.5 2V.5a.5.5 0 0 0-1 0V2h-1V.5a.5.5 0 0 0-1 0V2h-1V.5a.5.5 0 0 0-1 0V2h-1V.5zm1 4.5h3A1.5 1.5 0 0 1 11 6.5v3A1.5 1.5 0 0 1 9.5 11h-3A1.5 1.5 0 0 1 5 9.5v-3A1.5 1.5 0 0 1 6.5 5z"
        />
      </svg>
      <div>
        <h4 class="fw-bold mb-0">Entwicklung</h4>
        <p>
          Ksrg-Connect wird stets weiter entwickelt und nutzt die neusten
          Web-Technologien.
        </p>
      </div>
    </div>
    <div class="col d-flex align-items-start">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="1.75em"
        height="1.75em"
        fill="currentColor"
        class="bi text-muted flex-shrink-0 me-3"
        viewBox="0 0 16 16"
      >
        <path
          d="M5.072.56C6.157.265 7.31 0 8 0s1.843.265 2.928.56c1.11.3 2.229.655 2.887.87a1.54 1.54 0 0 1 1.044 1.262c.596 4.477-.787 7.795-2.465 9.99a11.775 11.775 0 0 1-2.517 2.453 7.159 7.159 0 0 1-1.048.625c-.28.132-.581.24-.829.24s-.548-.108-.829-.24a7.158 7.158 0 0 1-1.048-.625 11.777 11.777 0 0 1-2.517-2.453C1.928 10.487.545 7.169 1.141 2.692A1.54 1.54 0 0 1 2.185 1.43 62.456 62.456 0 0 1 5.072.56z"
        />
      </svg>
      <div>
        <h4 class="fw-bold mb-0">Datenschutz</h4>
        <p>
          Diese Website nutzt die höchsten Datenschutzstandards und verarbeitet
          alle Daten zu 100% 🇩🇪 in Deutschland.
        </p>
      </div>
    </div>
  </div>
</div>
