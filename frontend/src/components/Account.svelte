<script>
  import { onMount } from "svelte";
  import { navigate } from "svelte-navigator";
  import { loggedIn } from "../store";

  let username = "";
  let firstName = "";
  let lastName = "";
  let schoolClass = "";
  let age = 0;
  let description = "";
  let contact = "";
  let password = "";

  let result = null;
  let data = null;

  onMount(async () => {
    const res = await fetch(API_URL + "api/profile", {
      method: "GET",
      credentials: "include",
    });

    const json = await res.json();
    data = JSON.parse(JSON.stringify(json));

    username = data.username;
    firstName = data.firstName;
    lastName = data.lastName;
    schoolClass = data.schoolClass;
    age = data.age;
    description = data.description;
    contact = data.contact;
  });

  async function updateUser(e) {
    e.preventDefault();
    const res = await fetch(API_URL + "api/update/user", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        firstName,
        lastName,
        schoolClass,
        age,
        description,
        contact,
      }),
    });

    const json = await res.json();
    result = JSON.parse(JSON.stringify(json));
    console.log(result)
  }

  async function updatePassword(e) {
    e.preventDefault();
    const res = await fetch(API_URL + "api/update/password", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        password,
      }),
    });

    const json = await res.json();
    result = JSON.parse(JSON.stringify(json));
    console.log(result);
  }

  async function updateUsername(e) {
    e.preventDefault();
    const res = await fetch(API_URL + "api/update/username", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        username,
      }),
    });

    const json = await res.json();
    result = JSON.parse(JSON.stringify(json));
    console.log(result);
  }

  async function deleteUser(e) {
    e.preventDefault();

    const res = await fetch(API_URL + "api/delete", {
      method: "POST",
      credentials: "include",
    });

    const json = await res.json();
    result = JSON.parse(JSON.stringify(json));

    if (result.success) {
      loggedIn.set(false);
      navigate("/");
    }
  }
</script>

{#if result}
  {#if result.success}
    <div class="alert alert-success" role="alert">
      {result.msg}
    </div>
  {:else}
    <div class="alert alert-danger" role="alert">
      {result.msg}
    </div>
  {/if}
{/if}
<form>
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
    <div id="descriptionHelp" class="form-text">
      Tell us something about you.
    </div>
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
  <button type="submit" class="btn btn-primary" on:click={updateUser}
    >Update information</button
  >
</form>

<hr />

<form>
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
  <button type="submit" class="btn btn-primary" on:click={updatePassword}
    >Update password</button
  >
</form>

<hr />

<form>
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
      A change of your username could result in problems. Please be really sure
      about it.
    </div>
  </div>
  <button type="submit" class="btn btn-secondary" on:click={updateUsername}
    >Update username</button
  >
</form>

<hr />

<form>
  <div class="mb-3">
    If you delete your account all data will be forever lost. Are you really
    sure?
  </div>

  <button type="button" class="btn btn-danger" on:click={deleteUser}>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="16"
      height="16"
      fill="currentColor"
      class="bi bi-trash3-fill"
      viewBox="0 0 16 16"
    >
      <path
        d="M11 1.5v1h3.5a.5.5 0 0 1 0 1h-.538l-.853 10.66A2 2 0 0 1 11.115 16h-6.23a2 2 0 0 1-1.994-1.84L2.038 3.5H1.5a.5.5 0 0 1 0-1H5v-1A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5Zm-5 0v1h4v-1a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5ZM4.5 5.029l.5 8.5a.5.5 0 1 0 .998-.06l-.5-8.5a.5.5 0 1 0-.998.06Zm6.53-.528a.5.5 0 0 0-.528.47l-.5 8.5a.5.5 0 0 0 .998.058l.5-8.5a.5.5 0 0 0-.47-.528ZM8 4.5a.5.5 0 0 0-.5.5v8.5a.5.5 0 0 0 1 0V5a.5.5 0 0 0-.5-.5Z"
      />
    </svg>
    Delete account</button
  >
</form>
