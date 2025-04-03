import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';

export const authGuard: CanActivateFn = (route, state) => {
  const router = inject(Router);
  console.log(route, state);

  const token = localStorage.getItem('token');
  if (token) {
    return true; // Allow access
  } else {
    router.navigate(['/login']);
    return false; // Prevent access to the route
  }
};
