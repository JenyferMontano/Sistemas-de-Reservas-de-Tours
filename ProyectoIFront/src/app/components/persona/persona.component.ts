import { Component } from '@angular/core';
import { Persona } from '../../models/persona';
import { PersonaService } from '../../services/persona.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { UsuarioService } from '../../services/usuario.service';

@Component({
  selector: 'app-persona',
  imports: [CommonModule, FormsModule],
  templateUrl: './persona.component.html',
  styleUrl: './persona.component.css',
  providers: [PersonaService]
})
export class PersonaComponent {
    public status:number
  public persona:Persona
  private token:any

  constructor(
    private usuarioService:UsuarioService,
    private personaService:PersonaService
  ){
    this.status=-1
    this.persona=new Persona(0, '', '', '', new Date(), '', '')
    
  }

    // Getter para mostrar la fecha en formato yyyy-MM-dd
  get fechaNacString(): string {
    if (!this.persona.fechaNac) return '';
    return this.persona.fechaNac.toISOString().substring(0, 10);
  }

  // Setter para actualizar la fecha a partir del input
  set fechaNacString(value: string) {
    this.persona.fechaNac = new Date(value);
  }

  crearPersona() {
    this.token=this.usuarioService.getToken();
    if (!this.token) {
      alert('Token de autenticación no definido.');
      return;
    }
    this.personaService.crearPersona(this.persona,this.token).subscribe({
      next:(response:any)=>{
        alert('Persona creada con éxito!!');
        console.log('Respuesta:', response);
      },
      error:(err:Error)=>{
        console.error('Error al crear persona:', err);
      alert('Error al crear persona.');
      }
    })


}

 


  /*
   nuevaPersona: Persona = new Persona(0, '', '', '', new Date(), '', '');
     get fechaNacString(): string {
    if (!this.nuevaPersona.fechaNac) return '';
    const d = new Date(this.nuevaPersona.fechaNac);
    return d.toISOString().substring(0, 10);
  }

  set fechaNacString(value: string) {
    this.nuevaPersona.fechaNac = new Date(value);
  }
  token: string = ''; 

  constructor(private personaService: PersonaService) {}

  crearPersona() {
  const token = sessionStorage.getItem('token');
  if (!token) {
    alert('Token de autenticación no definido.');
    return;
  }

  this.personaService.crearPersona(this.nuevaPersona, token).subscribe({
    next: (response) => {
      alert('Persona creada con éxito.');
      // Aquí limpiar formulario o recargar datos
    },
    error: (err) => {
      console.error('Error al crear persona:', err);
      alert('Error al crear persona.');
    }
  });
}
  */
}
