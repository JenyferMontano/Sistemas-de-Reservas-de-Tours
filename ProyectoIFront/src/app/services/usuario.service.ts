
import {Injectable} from "@angular/core"
import {server} from "./global"
import {HttpClient,HttpHeaders} from "@angular/common/http"
import { Observable } from "rxjs"
import { Usuario } from "../models/usuario"
import { LoginR } from '../models/loginR';


@Injectable({providedIn:'root'}) export class UsuarioService{
    private url:string
    constructor(
        private _http:HttpClient
    ){
        this.url=server.url
    }

     /**
   * Login de usuario
   * @param loginData Objeto con email y password
   */
  /*

    login(loginData: LoginR): Observable<any> {
        let userJSON=JSON.stringify(loginData)
        let headers =new HttpHeaders().set('Content-Type','application/json')
        let options={
            headers
        }
        return this._http.post(this.url+'login',userJSON,options)
  }
    getIdentity(){
        let identity=sessionStorage.getItem('identity')
        if(identity){
            return JSON.parse(identity)
        }
        return null
    }
    getToken(){
        return sessionStorage.getItem('token')
    }

    */
     login(loginData: LoginR): Observable<any> {
        let userJSON=JSON.stringify(loginData)
        let headers =new HttpHeaders().set('Content-Type','application/json')
        let options={
            headers
        }
        return this._http.post(this.url+'login',userJSON,options)
  }

     getIdentity(){
        let identity=sessionStorage.getItem('identity')
        if(identity){
            return JSON.parse(identity)
        }
        return null
    }

  /**
   * Obtener token almacenado en sessionStorage
   */
  getToken(){
        return sessionStorage.getItem('token')
    }
  
    //POST
    //crearUsuario(usuario: Usuario): Observable<any> {return this._http.post(this.url + 'usuarios', usuario);}
    
crearUsuario(usuario: Usuario, token: any): Observable<any> {
  const accessToken = 'bearer ' + token;
  const headers = new HttpHeaders()
    .set('Content-Type', 'application/json')
    .set('Authorization', accessToken);

  const options = { headers };
  const data = JSON.stringify(usuario);

  return this._http.post(this.url + 'usuario', data, options);
}


}