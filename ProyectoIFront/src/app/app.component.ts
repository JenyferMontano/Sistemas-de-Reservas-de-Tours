import { Component } from '@angular/core';
import { RouterLink, RouterOutlet,  } from '@angular/router';
import { PersonaService } from "./services/persona.service"
import { UsuarioService } from './services/usuario.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'ProyectoIFront';
  //public personas:any;
  //private checkCategories;
  private checkIdentity;
  public identity:any;

  constructor(
    private personaService:PersonaService,
    private usuarioService:UsuarioService
  
  ){
    this.checkIdentity=setInterval(()=>{
      this.identity=usuarioService.getIdentity();
    },500)
  }
}

