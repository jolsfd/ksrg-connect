<script>
  export let firstName = "";
  export let lastName = "";
  export let schoolClass = "";
  export let age = 0;
  export let description = "";
  export let contact = "";
  export let username = "";

  export let showUsername = false;

</script>

<div class="card border-dark mt-2">
  <div class="card-body">
    {#if showUsername}
      <h4 class="card-title">{username}</h4>
    {/if}

    <h5 class="card-title">{firstName}, {lastName}</h5>
    <h6 class="card-subtitle mb-2 text-muted">
      {age} years old; class: {schoolClass}
    </h6>
    <p class="card-text">{description}</p>
    <details>
      <summary>Contact information</summary>
      <p>{contact}</p>
    </details>
  </div>
</div>
