<script>
    //export let username = '';
    export let firstName = '';
    export let lastName = '';
    export let schoolClass = '';
    export let age = 0;
    export let description = '';
    export let contact = '';
</script>

<div class="card">
    <div class="card-body">
      <h5 class="card-title">{firstName}, {lastName}</h5>
      <h6 class="card-subtitle mb-2 text-muted">{age} years old; {schoolClass}</h6>
      <p class="card-text">{description}</p>
      <details>
          <summary>Contact information</summary>
          <p>{contact}</p>
      </details>
    </div>
</div>