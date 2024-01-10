<script setup>
</script>

<template>
  <h1>НЕДВИЖИМОСТЬ</h1>
    <p>Добро пожаловать!<br>Мы занимаемся продажей недвижимости!
    Выполняем свою работу быстро и качественно!</p>
 
    <div id="carouselExampleControls" class="carousel slide" data-bs-ride="carousel">
        <div class="carousel-inner">
          <div class="carousel-item active">
            <img src="assets/img/main.jpg" class="d-block w-100" alt="...">
          </div>
          <div class="carousel-item">
            <img src="assets/img/t1.jpeg" class="d-block w-100" alt="...">
          </div>
          <div class="carousel-item">
            <img src="assets/img/t2.jpeg" class="d-block w-100" alt="...">
          </div>
        </div>
        <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleControls" data-bs-slide="prev">
          <span class="carousel-control-prev-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Previous</span>
        </button>
        <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleControls" data-bs-slide="next">
          <span class="carousel-control-next-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Next</span>
        </button>
      </div>
  
</template>

<script>
export default {
  data() {
    return {
    }
  },
}
</script>
