import {Injectable} from "@angular/core"
import {server} from "./global"
import {HttpClient,HttpHeaders} from "@angular/common/http"
import { Observable } from "rxjs"
import { Persona } from "../models/persona"

@Injectable({
    providedIn:'root'
})export class PersonaService{
    public url:string
    private accessToken:string
    constructor(private _http:HttpClient){
        this.url=server.url
        this.accessToken=""
    }
    getPersonas():Observable<any>{
       const headers=new HttpHeaders().set('Content-Type','application/json')       
        const options={
            headers
        }
        return this._http.get(this.url+'persona/',options)
    }

       crearPersona(persona:Persona,token:any):Observable<Persona>{
        this.accessToken="Bearer "+token
        let headers=new HttpHeaders().set('Content-Type','application/json').set('Authorization',this.accessToken)
        let options={
            headers
        }
          const body = {
    id_persona: persona.idPersona,
    nombre: persona.nombre,
    apellido_1: persona.apellido_1,
    apellido_2: persona.apellido_2,
    fecha_nac: persona.fechaNac.toISOString(),
    telefono: persona.telefono,
    correo: persona.correo
  };
        return this._http.post<Persona>(this.url+'persona/',body,options)

    }
        
/*
    crearPersona(persona: Persona, token: string): Observable<any> {
  this.accessToken = "Bearer " + token;

  const headers = new HttpHeaders()
    .set('Content-Type', 'application/json')
    .set('Authorization', this.accessToken);

  const body = {
    id_persona: persona.idPersona,
    nombre: persona.nombre,
    apellido_1: persona.apellido_1,
    apellido_2: persona.apellido_2,
    fecha_nac: persona.fechaNac.toISOString(),
    telefono: persona.telefono,
    correo: persona.correo
  };

  return this._http.post(this.url + 'persona/', body, { headers });
}

*/
 
}