<script>
  import { navigate } from "svelte-navigator";
  import Result from "./Result.svelte";

  // Submit values
  let username = "";
  let firstName = "";
  let lastName = "";
  let schoolClass = "";
  let age = 0;
  let description = "";
  let contact = "";
  let password = "";
  let authPassword = "";

  let result = null;

  async function signUp() {
    const res = await fetch(API_URL + "api/signup", {
      method: "POST",
      body: JSON.stringify({
        username,
        firstName,
        lastName,
        schoolClass,
        age,
        description,
        contact,
        password,
        authPassword,
      }),
    });

    const json = await res.json();

    result = JSON.parse(JSON.stringify(json));

    if (result.success) {
      navigate("/signin");
    }
  }

  function nav() {
    navigate("/about");
  }
</script>

<h2 class="display-5 fw-bold text-center">Sign Up</h2>

<Result {result} />

<div class="mb-3">
  <label for="inputUsername" class="form-label">Username</label>
  <input
    type="text"
    class="form-control"
    id="inputUsername"
    aria-describedby="usernameHelp"
    bind:value={username}
  />
  <div id="usernameHelp" class="form-text">
    Make sure your username is free of special characters, spaces and emojis.
  </div>
</div>

<div class="mb-3">
  <label for="inputFirstName">First Name</label>
  <input
    type="text"
    aria-label="First name"
    class="form-control"
    id="inputFirstName"
    bind:value={firstName}
  />
  <label for="inputLastName">Last Name</label>
  <input
    type="text"
    aria-label="Last name"
    class="form-control"
    id="inputLastName"
    bind:value={lastName}
  />
</div>

<label for="inputSchoolClass">Select your school class</label>
<select
  bind:value={schoolClass}
  class="form-select"
  aria-label="Select your school Class"
  id="inputSchoolClass"
>
  <!-- <option selected>Select your school class</option> -->
  <option value="12">12</option>
  <option value="11">11</option>
  <option value="10a">10a</option>
  <option value="10b">10b</option>
  <option value="10c">10c</option>
  <option value="9a">9a</option>
  <option value="9b">9b</option>
  <option value="9c">9c</option>
  <option value="9d">9d</option>
  <option value="8a">8a</option>
  <option value="8b">8b</option>
  <option value="8c">8c</option>
  <option value="other">other</option>
</select>

<div class="mb-3">
  <label for="inputAge" class="form-label">Age</label>
  <input type="number" class="form-control" id="inputAge" bind:value={age} />
</div>

<div class="mb-3">
  <label for="textareaDescription" class="form-label">Description</label>
  <textarea
    class="form-control"
    id="textareaDescription"
    rows="3"
    aria-describedby="descriptionHelp"
    bind:value={description}
  />
  <div id="descriptionHelp" class="form-text">Tell us something about you.</div>
</div>

<div class="mb-3">
  <label for="inputContact" class="form-label">Contact information</label>
  <input
    type="text"
    class="form-control"
    id="inputContact"
    aria-describedby="contactHelp"
    bind:value={contact}
  />
  <div id="contactHelp" class="form-text">
    For example your lernsax e-mail or instagram account.
  </div>
</div>

<div class="mb-3">
  <label for="inputPassword" class="form-label">Password</label>
  <input
    type="password"
    class="form-control"
    id="inputPassword"
    aria-describedby="passwordHelp"
    bind:value={password}
  />
  <div id="passwordHelp" class="form-text">
    Your password must be 8-40 characters long and must not contain emojis.
  </div>
</div>

<div class="mb-3">
  <label for="inputAuthPassword" class="form-label"
    >Authentication password</label
  >
  <input
    type="text"
    class="form-control"
    id="inputAuthPassword"
    aria-describedby="authPasswordHelp"
    bind:value={authPassword}
  />
  <div id="authPasswordHelp" class="form-text">
    You can get the authentication password by certain people.
  </div>
</div>

<!-- <div class="mb-3 form-check">
  <input type="checkbox" class="form-check-input" id="exampleCheck1" />
  <label class="form-check-label" for="exampleCheck1"
    >I agree to the terms of use</label
  >
</div> -->
<button type="submit" class="btn btn-primary" on:click={signUp}>Sign up</button>
<button class="btn btn-secondary" on:click={nav}>Help</button>

<hr class="my-4" />
<small class="text-muted"
  >By clicking Sign up, you agree to the terms of use.</small
>