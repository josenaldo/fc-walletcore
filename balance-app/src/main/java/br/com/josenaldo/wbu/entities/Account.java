package br.com.josenaldo.wbu.entities;

import jakarta.persistence.*;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.PositiveOrZero;
import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;

import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "accounts")
public class Account {

    @EmbeddedId
    @AttributeOverride(name = "value", column = @Column(name = "id", nullable = false))
    private EntityId id;

    @CreationTimestamp
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @UpdateTimestamp
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    @NotNull
    @Embedded
    @AttributeOverride(name = "value", column = @Column(name = "client_id", nullable = false))
    private EntityId clientId;

    @NotNull
    @PositiveOrZero
    @Column(name = "balance", nullable = false)
    private BigDecimal balance;

    public Account() {
        this.id = new EntityId();
    }

    public Account(String id, String clientId) {
        this.id = new EntityId(id);
        this.createdAt = LocalDateTime.now();
        this.updatedAt = LocalDateTime.now();
        this.clientId = new EntityId(clientId);
        this.balance = BigDecimal.ZERO;
    }

    public Account(EntityId id, LocalDateTime createdAt, LocalDateTime updatedAt, EntityId clientId, BigDecimal balance) {
        this.id = id;
        this.createdAt = createdAt;
        this.updatedAt = updatedAt;
        this.clientId = clientId;
        this.balance = balance;
    }

    public EntityId getId() {
        return id;
    }

    public void setId(EntityId id) {
        this.id = id;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt;
    }

    public LocalDateTime getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt(LocalDateTime updatedAt) {
        this.updatedAt = updatedAt;
    }

    public @NotNull EntityId getClientId() {
        return clientId;
    }

    public void setClientId(@NotNull EntityId clientId) {
        this.clientId = clientId;
    }

    @NotNull
    @PositiveOrZero
    public BigDecimal getBalance() {
        return balance;
    }

    public void setBalance(@NotNull @PositiveOrZero BigDecimal balance) {
        this.balance = balance;
    }
}