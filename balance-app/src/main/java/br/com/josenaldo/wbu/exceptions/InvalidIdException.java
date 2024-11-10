package br.com.josenaldo.wbu.exceptions;

public class InvalidIdException extends RuntimeException {
    public InvalidIdException() {
        super("Invalid entity id. The value must be a ULID.");
    }
}
