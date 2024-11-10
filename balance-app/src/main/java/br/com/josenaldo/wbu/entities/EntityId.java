package br.com.josenaldo.wbu.entities;

import br.com.josenaldo.wbu.exceptions.InvalidIdException;
import com.github.f4b6a3.ulid.Ulid;
import com.github.f4b6a3.ulid.UlidCreator;
import jakarta.persistence.Embeddable;
import jakarta.validation.constraints.NotNull;

import java.io.Serial;
import java.io.Serializable;
import java.util.Objects;

@Embeddable
public class EntityId implements Serializable {

    @Serial
    private static final long serialVersionUID = 1L;

    @NotNull
    private final String value;

    public EntityId() {
        Ulid ulid = UlidCreator.getUlid();
        this.value = ulid.toString();
    }

    public EntityId(String value) {
        try {
            Ulid ulid = Ulid.from(value);
            this.value = ulid.toString();
        } catch (IllegalArgumentException e) {
            throw new InvalidIdException();
        }
    }

    public String getValue() {
        return value;
    }

    public String toString() {
        return value;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        EntityId entityId = (EntityId) o;
        return Objects.equals(value, entityId.value);
    }

    @Override
    public int hashCode() {
        return Objects.hashCode(value);
    }
}
