package br.com.josenaldo.wbu.listeners;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.ZonedDateTime;

public class KafkaMessage {

    @JsonProperty("Name")
    private String name;

    @JsonProperty("Payload")
    private Payload payload;

    @JsonProperty("CreatedAt")
    private ZonedDateTime createdAt;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Payload getPayload() {
        return payload;
    }

    public void setPayload(Payload payload) {
        this.payload = payload;
    }

    public ZonedDateTime getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(ZonedDateTime createdAt) {
        this.createdAt = createdAt;
    }

    @Override
    public String toString() {
        return "KafkaMessage{" +
                "name='" + name + '\'' +
                ", payload=" + payload +
                ", createdAt=" + createdAt +
                '}';
    }
}
