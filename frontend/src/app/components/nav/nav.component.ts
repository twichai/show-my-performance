import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-nav',
  imports: [],
  templateUrl: './nav.component.html',
  styleUrl: './nav.component.css',
})
export class NavComponent {
  router = inject(Router);
  goToCreatePost() {
    this.router.navigate(['/create-post']);
  }

  logout() {
    localStorage.removeItem('token'); // Remove user token
    this.router.navigate(['/login']); // Redirect to login page
  }
}
