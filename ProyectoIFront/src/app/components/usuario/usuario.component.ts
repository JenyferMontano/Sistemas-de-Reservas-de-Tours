import { Component } from '@angular/core';
import { UsuarioService } from '../../services/usuario.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Usuario } from '../../models/usuario';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-usuario',
  imports: [CommonModule, FormsModule],
  templateUrl: './usuario.component.html',
  styleUrl: './usuario.component.css',
  providers: [UsuarioService]
})
export class UsuarioComponent {
 public status: string = ''; // 'success' | 'error' | 'unauthorized'
  public usuario: Usuario;

  constructor(
    private _usuarioService: UsuarioService,
    private _router: Router,
    private _routes: ActivatedRoute
  ) {
     //this.status = -1;
      this.usuario=new Usuario("","","",0)
  }

  onSubmit() {
    const token = this._usuarioService.getToken();
    if (!token) {
      this.status = 'unauthorized';
      return;
    }

    this._usuarioService.crearUsuario(this.usuario, token).subscribe({
      next: (res) => {
        console.log('Usuario creado:', res);
        this.status = 'success';
        this._router.navigate(['']); 
      },
      error: (err) => {
        console.error('Error al crear usuario:', err);
        this.status = 'error';
      }
    });
  }

}
