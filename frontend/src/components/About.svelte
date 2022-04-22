<script>
  import { Link } from "svelte-navigator";
</script>

<div class="px-4 my-5">
  <h4 class="display-5 fw-bold text-center">Quickstart</h4>
  <div class="col-lg-9 mx-auto">
    <h6 class="display-6 fw-bold">Registrierung</h6>
    <p class="lead mb-4">
      Registriere dich über die <Link to="/signup" primary={false} role="link"
        >SignUp</Link
      > Seite. Du musst einen Benutzernamen und ein Passwort, welches mindestens
      8 Zeichen lang sein muss, für deinen Account hinterlegen. Danach gibst du deine
      persönlichen Daten ein. Diese umfassen deinen Klarnamen, dein Alter, eine kurze
      interessante Beschreibung von dir und deine Kontaktdaten. Du kannst für deine
      Kontaktdaten zum Beispiel deinen Instagram Account oder deine Lernsax E-Mail
      hinterlegen. Um deine Registrierung abschließen zu können musst du das Authentifizierungspasswort
      eingeben. Dieses kannst du über deinen Klassensprecher erhalten. Wenn deine
      Registrierung erfolgreich war, kannst du dich anschließend auf der <Link
        to="/signin"
        primary={false}
        role="link">SignIn</Link
      > Seite mit deinem Benutzernamen und deinem Passwort einloggen.
    </p>

    <h6 class="display-6 fw-bold">Nutzung</h6>
    <p class="lead" mb-4>
      Nach erfolgreichem Einloggen kannst du auf der <Link
        to="/explore"
        primary={false}
        role="link">Explore</Link
      > Seite nach Personen stöbern. Du kannst auch direkt nach bestimmten Details
      wie z.B. Alter, Klasse und Namen in der Suchleiste filtern. Über die Schaltfläche
      Contact Information kannst du dir die Kontaktdaten dieser Person ansehen. Auf
      der <Link to="/account" primary={false} role="link">Account</Link> Seite kannst
      du deinen Account bearbeiten oder löschen.
    </p>
  </div>
</div>

<div class="px-4 my-5 text-center">
  <h4 class="display-5 fw-bold">Häufig gestellte Fragen</h4>
  <div class="col-lg-9 mx-auto">
    <div class="accordion" id="accordionExample">
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingOne">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseOne"
            aria-expanded="true"
            aria-controls="collapseOne"
          >
            #1 Warum KSRG-Connect?
          </button>
        </h2>
        <div
          id="collapseOne"
          class="accordion-collapse collapse"
          aria-labelledby="headingOne"
          data-bs-parent="#accordionExample"
        >
          <div class="accordion-body">
            Wir haben KSRG-Connect entwickelt damit <strong>du</strong> dich aktiv
            mit neuen Menschen vernetzen kannst.
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingTwo">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseTwo"
            aria-expanded="false"
            aria-controls="collapseTwo"
          >
            #2 Wann wird Feature X kommen?
          </button>
        </h2>
        <div
          id="collapseTwo"
          class="accordion-collapse collapse"
          aria-labelledby="headingTwo"
          data-bs-parent="#accordionExample"
        >
          <div class="accordion-body">
            KSRG-Connect wurde mit einem minimalen Funktionsumfang gestartet,
            den wir für notwendig halten, um den Benutzern eine angenehme
            Webxperience bieten zu können. Zukünftige Features können allerdings
            eine gewisse Zeit benötigen, um veröffenticht zu werden, da
            KSRG-Connect ein Hobbyprojekt ist und interne Entwicklungssysteme
            möglicherweise angepasst oder umgebaut werden müssen damit neue
            Features bereit gestellt werden können.
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingThree">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseThree"
            aria-expanded="false"
            aria-controls="collapseThree"
          >
            #3 Wie kann ich KSRG-Connect unterstützen?
          </button>
        </h2>
        <div
          id="collapseThree"
          class="accordion-collapse collapse"
          aria-labelledby="headingThree"
          data-bs-parent="#accordionExample"
        >
          <div class="accordion-body">
            Du kannst uns unterstützen indem du uns entweder deinen
            Mitlernendenweiter empfiehlst oder aktive Entwicklungsarbeit am
            Quellcode leistet.
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="headingFour">
          <button
            class="accordion-button collapsed"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#collapseFour"
            aria-expanded="false"
            aria-controls="collapseFour"
          >
            #4 Wie lange wird es KSRG-Connect geben?
          </button>
        </h2>
        <div
          id="collapseFour"
          class="accordion-collapse collapse"
          aria-labelledby="headingThree"
          data-bs-parent="#accordionExample"
        >
          <div class="accordion-body">
            Wir sehen uns gezwungen bei Verstößen z.B. Beleidigungen, Hassrede
            uvm. KSRG-Connect umgehend abzuschalten. Desweiteren werden Mittel
            benötigt, um Server und Domain zu finanzieren. Diese Mittel sind bis
            Herbst 2022 gesichert danach sind wir auf Unterstützung angewiesen.
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
