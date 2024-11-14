package br.com.josenaldo.wbu.listeners;

import br.com.josenaldo.wbu.services.AccountService;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;

@Slf4j
@Component
public class BalanceUpdatedKafkaListener {
    private AccountService accountService;
    private final ObjectMapper objectMapper;

    public BalanceUpdatedKafkaListener(AccountService accountService) {
        this.accountService = accountService;
        this.objectMapper = new ObjectMapper();
        this.objectMapper.registerModule(new JavaTimeModule());
    }

    @KafkaListener(topics = "balances", groupId = "balance-updater")
    public void listenGroupFoo(@Payload String message) {
        try {
            KafkaMessage kafkaMessage = objectMapper.readValue(message, KafkaMessage.class);
            log.info("Received Message in group foo: {}", kafkaMessage);

            String accountIdFrom = kafkaMessage.getPayload().getAccountIdFrom();
            String accountIdTo = kafkaMessage.getPayload().getAccountIdTo();
            BigDecimal balanceFrom = kafkaMessage.getPayload().getBalanceFrom();
            BigDecimal balanceTo = kafkaMessage.getPayload().getBalanceTo();

            accountService.updateBalances(
                    accountIdFrom,
                    accountIdTo,
                    balanceFrom,
                    balanceTo
            );

            log.info("Balances updated successfully updated from {} to {}", accountIdFrom, accountIdTo);
        } catch (Exception e) {
            log.error("Error processing message", e);
        }

    }
}
