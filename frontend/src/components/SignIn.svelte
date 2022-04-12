<script>
  import { loggedIn } from "../store";
  import { navigate } from "svelte-navigator";
  import Result from "./Result.svelte";

  let username = "";
  let password = "";

  let result = null;

  async function signIn() {
    const res = await fetch(API_URL + "api/signin", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        username,
        password,
      }),
    });

    const json = await res.json();
    result = JSON.parse(JSON.stringify(json));
    console.log(result);

    if (result.success) {
      loggedIn.set(true);
      navigate("/explore");
    }
  }

  function handleKeydown(e){
    if (e.keyCode === 13) {
      signIn()
    }
  }
</script>

<svelte:window on:keydown={handleKeydown}/>

<h2 class="display-5 fw-bold text-center">Sign In</h2>

<Result result={result} />

<div class="mb-3">
  <label for="inputUsername" class="form-label">Username</label>
  <input
    type="text"
    class="form-control"
    id="inputUsername"
    bind:value={username}
  />
</div>

<div class="mb-3">
  <label for="inputPassword" class="form-label">Password</label>
  <input
    type="password"
    class="form-control"
    id="inputPassword"
    bind:value={password}
  />
</div>

<button type="submit" class="btn btn-primary" on:click={signIn}>Sign in</button>
