import { Component } from '@angular/core';
import { Usuario } from '../../models/usuario';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { UsuarioService } from '../../services/usuario.service';
import { LoginR } from '../../models/loginR';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-login',
  imports: [CommonModule, FormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
  providers:[UsuarioService]
})

export class LoginComponent {
  public status: number;
 // public usuario: Usuario;
  public loginData: LoginR;

  constructor(
    private _usuarioService: UsuarioService,
    private _router: Router,
    private _routes: ActivatedRoute
  ) {
     this.status = -1;
    this.loginData = { email: '', password: '' };;
  }

    onSubmit() {
    this._usuarioService.login(this.loginData).subscribe({
      next: (response: any) => {
        if (response.access_token && response.user) {
          sessionStorage.setItem('token', response.access_token);
          sessionStorage.setItem('identity', JSON.stringify(response.user));
          this._router.navigate(['']);
          // Mostrar en consola username y rol QUITAR DESPUEES ESTO ES PARAR PRUEBAS
          console.log('Usuario logueado:');
          console.log('Username:', response.user.username);
          console.log('Rol:', response.user.role);
        } else {
          this.status = 0; // Login incorrecto
        }
      },
      error: (err) => {
        console.error('Error en login:', err);
        this.status = 1; // Error de servidor
      }
    });
  }
/*
  onSubmit(form: any) {
    this._usuarioService.login(this.loginData).subscribe({
      next: (response: any) => {
        if (response.token && response.user) {
          sessionStorage.setItem('token', response.token);
          sessionStorage.setItem('identity', JSON.stringify(response.user));
          this._router.navigate(['']);
      // Mostrar en consola username y rol
        console.log('Usuario logueado:');
        console.log('Username:', response.user.username);
        console.log('Rol:', response.user.role);
        } else {
          this.status = 0; // Login incorrecto
        }
      },
      error: (err) => {
        console.log(err);
        this.status = 1; // Error de servidor
      }
    });
  }*/
}
