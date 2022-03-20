<script>
  import { Router, Link, Route } from "svelte-navigator";
  import { loggedIn } from "./store";
  import SignUp from "./components/SignUp.svelte";
  import SignIn from "./components/SignIn.svelte";
  import Account from "./components/Account.svelte";
  import SignOut from "./components/SignOut.svelte";
  import Users from "./components/Users.svelte";
  import Home from "./components/Home.svelte";
  import About from "./components/About.svelte";

  export let name;

  console.log("Frontend started successfully", name, API_URL);

  let signedIn;

  loggedIn.subscribe((value) => {
    signedIn = value;
  });
</script>

<svg xmlns="http://www.w3.org/2000/svg" style="display: none;">
  <symbol id="github" viewBox="0 0 16 16">
    <title>GitHub</title>
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"
    />
  </symbol>
</svg>

<Router>
  <div class="wrapper">
    <nav class="navbar navbar-expand-md navbar-dark bg-dark">
      <div class="container-fluid">
        <Link to="/" class="navbar-brand">
          <img
            class="bi me-2"
            width="40"
            height="32"
            src="ksrg-connect_logo.svg"
            alt="Bootstrap"
          />
        </Link>
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon" />
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <Link to="account" class="nav-link active" aria-current="page"
                >Account <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="currentColor"
                  class="bi bi-gear"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="M8 4.754a3.246 3.246 0 1 0 0 6.492 3.246 3.246 0 0 0 0-6.492zM5.754 8a2.246 2.246 0 1 1 4.492 0 2.246 2.246 0 0 1-4.492 0z"
                  />
                  <path
                    d="M9.796 1.343c-.527-1.79-3.065-1.79-3.592 0l-.094.319a.873.873 0 0 1-1.255.52l-.292-.16c-1.64-.892-3.433.902-2.54 2.541l.159.292a.873.873 0 0 1-.52 1.255l-.319.094c-1.79.527-1.79 3.065 0 3.592l.319.094a.873.873 0 0 1 .52 1.255l-.16.292c-.892 1.64.901 3.434 2.541 2.54l.292-.159a.873.873 0 0 1 1.255.52l.094.319c.527 1.79 3.065 1.79 3.592 0l.094-.319a.873.873 0 0 1 1.255-.52l.292.16c1.64.893 3.434-.902 2.54-2.541l-.159-.292a.873.873 0 0 1 .52-1.255l.319-.094c1.79-.527 1.79-3.065 0-3.592l-.319-.094a.873.873 0 0 1-.52-1.255l.16-.292c.893-1.64-.902-3.433-2.541-2.54l-.292.159a.873.873 0 0 1-1.255-.52l-.094-.319zm-2.633.283c.246-.835 1.428-.835 1.674 0l.094.319a1.873 1.873 0 0 0 2.693 1.115l.291-.16c.764-.415 1.6.42 1.184 1.185l-.159.292a1.873 1.873 0 0 0 1.116 2.692l.318.094c.835.246.835 1.428 0 1.674l-.319.094a1.873 1.873 0 0 0-1.115 2.693l.16.291c.415.764-.42 1.6-1.185 1.184l-.291-.159a1.873 1.873 0 0 0-2.693 1.116l-.094.318c-.246.835-1.428.835-1.674 0l-.094-.319a1.873 1.873 0 0 0-2.692-1.115l-.292.16c-.764.415-1.6-.42-1.184-1.185l.159-.291A1.873 1.873 0 0 0 1.945 8.93l-.319-.094c-.835-.246-.835-1.428 0-1.674l.319-.094A1.873 1.873 0 0 0 3.06 4.377l-.16-.292c-.415-.764.42-1.6 1.185-1.184l.292.159a1.873 1.873 0 0 0 2.692-1.115l.094-.319z"
                  />
                </svg></Link
              >
            </li>
            <li class="nav-item">
              <Link to="explore" class="nav-link active" aria-current="page"
                >Explore <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="currentColor"
                  class="bi bi-heart"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z"
                  />
                </svg></Link
              >
            </li>
            <li class="nav-item">
              <Link to="about" class="nav-link active" aria-current="page"
                >About <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="currentColor"
                  class="bi bi-info-circle"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"
                  />
                  <path
                    d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"
                  />
                </svg></Link
              >
            </li>
          </ul>
          <div class="d-flex">
            <Link
              to="signup"
              class="btn btn-warning me-2"
              role="button"
              type="submit">Sign Up</Link
            >
            {#if signedIn}
              <Link
                to="signout"
                class="btn btn-outline-light"
                role="button"
                type="submit">Sign Out</Link
              >
            {:else}
              <Link
                to="signin"
                class="btn btn-outline-light"
                role="button"
                type="submit">Sign In</Link
              >
            {/if}
          </div>
        </div>
      </div>
    </nav>
    <div class="container">
      <Route path="/">
        <Home />
      </Route>
      <Route path="signin" primary={false}>
        <SignIn />
      </Route>
      <Route path="signup" primary={false}>
        <SignUp />
      </Route>
      <Route path="signout" primary={false}>
        <SignOut />
      </Route>
      <Route path="explore" primary={false}>
        <Users />
      </Route>
      <Route path="account" primary={false}>
        <Account />
      </Route>
      <Route path="about" primary={false}>
        <About />
      </Route>
    </div>
    <div class="container">
      <footer
        class="d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top"
      >
        <div class="col-md-4 d-flex align-items-center">
          <span class="text-muted">&copy; 2022 Jolsfd, Chemnitz</span>
        </div>

        <ul class="nav col-md-4 justify-content-end list-unstyled d-flex">
          <!-- <li class="ms-3">
            <a class="text-muted" href="github.com/jolsfd/ksrg-connect"
              ><svg class="bi" width="24" height="24"
                ><use xlink:href="#twitter" /></svg
              ></a
            >
          </li>
          <li class="ms-3">
            <a class="text-muted" href="google.com"
              ><svg class="bi" width="24" height="24"
                ><use xlink:href="#instagram" /></svg
              ></a
            >
          </li> -->
          <li class="ms-3">
            <a class="text-muted" href="https://github.com/jolsfd/ksrg-connect"
              ><svg class="bi" width="24" height="24"
                ><use xlink:href="#github" /></svg
              >
            </a>
          </li>
        </ul>
      </footer>
    </div>
  </div>
</Router>

<style>
  .wrapper {
    min-height: 100%;
    display: grid;
    grid-template-rows: auto 1fr auto;
  }

  nav {
    margin-bottom: 8px;
  }
</style>
